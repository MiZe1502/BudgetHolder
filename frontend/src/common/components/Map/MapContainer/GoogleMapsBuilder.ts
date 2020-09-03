import { PlacemarkBuilder } from "./PlacemarkBuilder";
import { MapsBuilder } from "./MapsBuilder";
import { MapItemData } from "../../ActionElements/MapActionElement/utils";


declare var google;

export class GoogleMapsBuilder implements MapsBuilder {
    map: any = null;
    placementBuilder?: PlacemarkBuilder;

    defaultZoom: number = 8;

    constructor(placementBuilder?: PlacemarkBuilder) {
        this.placementBuilder = placementBuilder;
    }

    private createMap(coords: Coordinates, container: HTMLDivElement) {
		this.map = new google.maps.Map(container, {
            zoomGmaps: this.defaultZoom,
			centerGmaps: coords,
		});
    }

    findMultipleAddressesAndCreateMap(data: MapItemData[], container: HTMLDivElement): void {

    };

    findSingleAddressAndCreateMap(data: MapItemData, container: HTMLDivElement): void {

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
