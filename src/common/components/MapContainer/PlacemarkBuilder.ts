import { Coordinates } from "./MapsBuilder";

export declare var ymaps;

export interface PlacemarkBuilder {
    createPlacemark: (coords: Coordinates) => any;
}

