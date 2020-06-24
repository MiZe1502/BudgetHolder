export interface MapsBuilder {
    container: HTMLDivElement;
    map: any;
    defaultZoom: number;
    findSingleAddressAndCreateMap: (address: string) => void;
}

export type Coordinates = [number, number];
