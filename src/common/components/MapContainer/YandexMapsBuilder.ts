import { PlacemarkBuilder } from "./PlacemarkBuilder";
import { MapsBuilder, Coordinates } from "./MapsBuilder";
import { MapItemData } from "../MapActionElement/utils";

declare var ymaps;

export class YandexMapsBuilder implements MapsBuilder {
    container: HTMLDivElement = null;
    map: any = null;
    placementBuilder?: PlacemarkBuilder;

    defaultZoom: number = 17;
    defaultCenter: Coordinates = [55.7, 37.5];
    defaultBaloonPreset: string = "islands#darkBlueDotIconWithCaption";

    constructor(container: any, placementBuilder?: PlacemarkBuilder) {
        this.container = container;
        this.placementBuilder = placementBuilder;
    }

    private processSingleAddressSearch(address: string): Promise<any> {
        return ymaps.geocode(address, {});
    }

    private getFoundedGeoObject(res: any, index: number) {
        return res.geoObjects.get(index);
    }

    private getFoundCoordinates(geoObject: any): Coordinates {
        return geoObject.geometry.getCoordinates()
    }

    private getVisibleArea(geoObject: any): Coordinates[] {
        return geoObject.geometry.getCoordinates()
    }

    private setGeoObjectProperties(geoObject): void {
        geoObject.options.set('preset', this.defaultBaloonPreset);
        geoObject.properties.set('iconCaption', geoObject.getAddressLine()); 
    }

    private addPlacementToMap(placement: any): void {
        this.map.geoObjects.add(placement);
    }

    findSingleAddressAndCreateMap(data: MapItemData): void {
        console.log(data)
        this.processSingleAddressSearch(data.address).then(res => {
            const geoObject = this.getFoundedGeoObject(res, 0);
            const coords = this.getFoundCoordinates(geoObject);
            const bounds = this.getVisibleArea(geoObject);

            this.createMap(coords);

            this.addGeoObjectAtMap(geoObject, bounds);
            this.setGeoObjectProperties(geoObject);

            const placement = this.placementBuilder.createPlacemark(coords, data);
            this.addPlacementToMap(placement)
        })
        .catch(err => console.log(err))
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


    findMultipleAddressesAndCreateMap(data: MapItemData[]): void {
        const promises: Promise<any>[] = [];
        data.forEach((data: MapItemData) => {
            promises.push(this.processSingleAddressSearch(data.address))
        });

        Promise.allSettled(promises).then((results) => {
            const pointsCollection = new ymaps.GeoObjectCollection();
            const points: Coordinates[] = [];

            results.forEach((res : PromiseFulfilledResult<any>, index: number) => {
                const currentRes = res.value;

                const geoObject = this.getFoundedGeoObject(currentRes, 0);
                const coords = this.getFoundCoordinates(geoObject);
                points.push(coords);

                const placement = this.placementBuilder.createPlacemark(coords, data[index]);
                pointsCollection.add(placement);
            })

            this.map = new ymaps.Map(this.container, {
                center: this.getCenter(points),
                zoom: this.defaultZoom,
                controls: ['zoomControl']
            });

            this.map.geoObjects.add(pointsCollection);
        })
        .catch(err => console.log(err))
    }

    private createMap(coords: Coordinates) {
        this.map = new ymaps.Map(this.container, {
            center: coords,
            zoom: this.defaultZoom
        });
    }

    private addGeoObjectAtMap(geoObject: any, bounds: any): void {
        this.map.geoObjects.add(geoObject);
        this.map.setBounds([bounds], {
            checkZoomRange: true,
        })
    }
}