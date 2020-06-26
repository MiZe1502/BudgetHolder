import { MapItemData } from "../MapActionElement/utils";

export interface MapsBuilder {
    container: HTMLDivElement;
    map: any;
    defaultZoom: number;
    findSingleAddressAndCreateMap: (data: MapItemData) => void;
    findMultipleAddressesAndCreateMap: (data: MapItemData[]) => void;
}

export type Coordinates = [number, number];
