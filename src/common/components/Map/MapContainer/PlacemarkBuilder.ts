import { Coordinates } from "./MapsBuilder";
import { MapItemData } from "../../ActionElements/MapActionElement/utils";

export declare var ymaps;

export interface PlacemarkBuilder {
    createPlacemark: (coords: Coordinates, data: MapItemData) => any;
}

