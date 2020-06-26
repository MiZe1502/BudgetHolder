import { PlacemarkBuilder } from "./PlacemarkBuilder";
import { MapsBuilder } from "./MapsBuilder";
import { MapItemData } from "../MapActionElement/utils";


declare var google;

export class GoogleMapsBuilder implements MapsBuilder {
    container: HTMLDivElement = null;
    map: any = null;
    placementBuilder?: PlacemarkBuilder;

    defaultZoom: number = 8;

    constructor(container: any, placementBuilder?: PlacemarkBuilder) {
        this.container = container;
        this.placementBuilder = placementBuilder;
    }

    private createMap(coords: Coordinates) {
		this.map = new google.maps.Map(this.container, {
            zoomGmaps: this.defaultZoom,
			centerGmaps: coords,
		});
    }

    findMultipleAddressesAndCreateMap(data: MapItemData[]): void {

    };

    findSingleAddressAndCreateMap(data: MapItemData): void {

		// const map = new google.maps.Map(container, {
        //     zoomGmaps,
		// 	centerGmaps,
		// });


        // const locations = [{
        //     latitude: 55.75361503443606,
        //     longitude: 37.620883000000006,
        //     name: 'Test point'
        // }];

        //const centerGmaps = {lat: -34.397, lng: 150.644};

        // const points = myMap.geoObjects;
        // locations.forEach(location => {
        //     points.add(new ymaps.Placemark(
        //         [location.latitude, location.longitude],
        //             {balloonContent: location.name},
        //             {
        //     }));
        // });

    }
}
