import { MapItemData } from "../../ActionElements/MapActionElement/utils";

export interface MapsBuilder {
    map: any;
    defaultZoom: number;
    findSingleAddressAndCreateMap: (data: MapItemData, container: HTMLDivElement) => void;
    findMultipleAddressesAndCreateMap: (data: MapItemData[], container: HTMLDivElement) => void;
}

export type Coordinates = [number, number];
