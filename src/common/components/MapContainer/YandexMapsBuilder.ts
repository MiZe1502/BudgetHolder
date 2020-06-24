import { PlacemarkBuilder } from "./PlacemarkBuilder";
import { MapsBuilder, Coordinates } from "./MapsBuilder";

declare var ymaps;

export class YandexMapsBuilder implements MapsBuilder {
    container: HTMLDivElement = null;
    map: any = null;
    placementBuilder?: PlacemarkBuilder;

    defaultZoom: number = 17;
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

    findSingleAddressAndCreateMap(address: string): void {
        this.processSingleAddressSearch(address).then(res => {
            const geoObject = this.getFoundedGeoObject(res, 0);
            const coords = this.getFoundCoordinates(geoObject);
            const bounds = this.getVisibleArea(geoObject);

            this.createMap(coords);

            this.addGeoObjectAtMap(geoObject, bounds);
            this.setGeoObjectProperties(geoObject);

            const placement = this.placementBuilder.createPlacemark(coords);
            this.addPlacementToMap(placement)
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
        this.map.setBounds(bounds, {
            checkZoomRange: true,
        })
    }

    
}