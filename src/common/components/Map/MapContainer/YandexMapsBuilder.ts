import {PlacemarkBuilder} from "./PlacemarkBuilder";
import {MapsBuilder, Coordinates} from "./MapsBuilder";
import {MapItemData} from "../../ActionElements/MapActionElement/utils";
import {Placemark} from "yandex-maps";

declare var ymaps;

export class YandexMapsBuilder implements MapsBuilder {
    map: any = null;
    placementBuilder?: PlacemarkBuilder;

    defaultZoom: number = 10;
    zoomedZoom: number = 17;
    defaultCenter: Coordinates = [55.76, 37.64];
    defaultBaloonPreset: string = "sislands#darkBlueDotIconWithCaption";

    singlePlacement: Placemark = null;
    isEditable: boolean = false;
    updateAddress: (newAddress: string) => void;

    constructor(placementBuilder?: PlacemarkBuilder, isEditable?: boolean, updateAddress?: (newAddress: string) => void) {
        this.placementBuilder = placementBuilder;
        this.isEditable = isEditable;
        this.updateAddress = updateAddress;
    }

    private processSingleAddressSearch(address: string): Promise<any> {
        console.log(address)

        return ymaps.geocode(address, {});
    }

    private processSingleCoordsSearch(coords: Coordinates): Promise<any> {
        return ymaps.geocode(coords);
    }

    private getFoundedGeoObject(res: any, index: number) {
        return res.geoObjects.get(index);
    }

    private getFoundCoordinates(geoObject: any): Coordinates {
        return geoObject.geometry.getCoordinates()
    }

    private getFoundAddress(geoObject: any): string {
        return geoObject.getAddressLine();
    }

    private getVisibleArea(geoObject: any): Coordinates[] {
        return geoObject.properties.get('boundedBy');
    }

    private setGeoObjectProperties(geoObject): void {
        geoObject.options.set('preset', this.defaultBaloonPreset);
        geoObject.properties.set('iconCaption', geoObject.getAddressLine());
    }

    private addPlacementToMap(placement: any): void {
        this.map.geoObjects.add(placement);
    }

    private findAddressByCoords(coords: Coordinates): void {
        this.processSingleCoordsSearch(coords).then(res => {
            this.removeAllDataFromMap();

            const geoObject = this.getFoundedGeoObject(res, 0);
            const address = this.getFoundAddress(geoObject);

            const bounds = this.getVisibleArea(geoObject);

            this.addGeoObjectAtMap(geoObject, bounds);
            this.setGeoObjectProperties(geoObject);

            this.updateAddress(address);
        });
    }

    makeMapEditable(data: MapItemData): void {
        const self = this;
        this.map.events.add("click", function (event: any) {
            const coords = event.get("coords");

            //move placement if exists
            if (self.singlePlacement) {
                // @ts-ignore
                self.singlePlacement.geometry.setCoordinates(coords);
            }
            // create it
            else {
                self.singlePlacement = self.placementBuilder.createPlacemark(coords, data);
                self.addPlacementToMap(self.singlePlacement)

                self.singlePlacement.events.add('dragend', function () {
                    // @ts-ignore
                    self.findAddressByCoords(self.singlePlacement.geometry.getCoordinates());
                });
            }
            self.findAddressByCoords(coords);
        });
    }

    findSingleAddressAndCreateMap(data: MapItemData, container: HTMLDivElement): void {
        if (!data || !container) {
            return;
        }

        if (!data.address) {
            this.createMap(this.defaultCenter, container, this.defaultZoom);
            if (this.isEditable) {
                this.makeMapEditable(data);
            }
        } else {
            this.processSingleAddressSearch(data.address).then(res => {
                console.log(res)

                const geoObject = this.getFoundedGeoObject(res, 0);
                const coords = this.getFoundCoordinates(geoObject);
                const bounds = this.getVisibleArea(geoObject);

                if (this.isMapExists()) {
                    this.removeAllDataFromMap();
                } else {
                    this.createMap(coords, container, this.zoomedZoom);
                }

                this.addGeoObjectAtMap(geoObject, bounds);
                this.setGeoObjectProperties(geoObject);

                this.singlePlacement = this.placementBuilder.createPlacemark(coords, data);
                this.addPlacementToMap(this.singlePlacement)

                if (this.isEditable) {
                    this.makeMapEditable(data);
                }
            })
                .catch(err => console.log(err))
        }
    }

    getCenter(points: Coordinates[]): Coordinates {
        let sumX = 0;
        let sumY = 0;
        points.forEach((point: Coordinates) => {
            sumX += point[0];
            sumY += point[1];
        })

        sumX = sumX / points.length;
        sumY = sumY / points.length;

        return [sumX, sumY];
    }

    private isMapExists(): boolean {
        return Boolean(this.map);
    }

    private removeAllDataFromMap(): void {
        this.map.geoObjects.removeAll();
    }

    findMultipleAddressesAndCreateMap(data: MapItemData[], container: HTMLDivElement): void {
        if (!data || !container) {
            return;
        }

        const promises: Promise<any>[] = [];
        data.forEach((data: MapItemData) => {
            promises.push(this.processSingleAddressSearch(data.address))
        });

        Promise.allSettled(promises).then((results) => {
            const pointsCollection = new ymaps.GeoObjectCollection();
            const points: Coordinates[] = [];

            results.forEach((res: PromiseFulfilledResult<any>, index: number) => {
                const currentRes = res.value;

                const geoObject = this.getFoundedGeoObject(currentRes, 0);
                const coords = this.getFoundCoordinates(geoObject);
                points.push(coords);

                const placement = this.placementBuilder.createPlacemark(coords, data[index]);
                pointsCollection.add(placement);
            })

            if (this.isMapExists()) {
                this.removeAllDataFromMap();
            } else {
                this.createMap(this.getCenter(points), container, this.zoomedZoom);
            }

            this.addPlacementToMap(pointsCollection);
            this.setVisibleAreaBasedOnMultiplePoints(pointsCollection);
        })
            .catch(err => console.log(err))
    }

    private setVisibleAreaBasedOnMultiplePoints(points: any) {
        this.map.setBounds(points.getBounds());
    }

    private createMap(coords: Coordinates, container: HTMLDivElement, zoom: number) {
        this.map = new ymaps.Map(container, {
            center: coords,
            zoom: zoom,
            controls: ['zoomControl']
        });
    }

    private addGeoObjectAtMap(geoObject: any, bounds: any): void {
        this.map.geoObjects.add(geoObject);
        this.map.setBounds(bounds, {
            checkZoomRange: true,
        })
    }
}