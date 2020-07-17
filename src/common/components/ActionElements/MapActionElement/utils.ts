
import { EntityType, ActionType } from "../../../../stores/utils";

export interface MapItemData {
    name: string;
    address: string;
    id: number;
    actionType: ActionType;
    entityType: EntityType;
}

export interface MapActionElementData {
    info: MapItemData[]
}