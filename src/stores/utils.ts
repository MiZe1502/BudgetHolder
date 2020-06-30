import { v4 as uuidv4 } from 'uuid';

export enum LoadingStatus {
    None = "None",
    Loading = "Loading",
    Finished = "Finished",
    Error = "Error",
};

export enum EntityType {
    Shop = "Shop",
}

export enum ActionType {
    Map = "Map",
    Edit = "Edit",
    Remove = "Remove",
    Add = "Add",
}

export const buildUniqueIdBasedOnEntityAndAction = (entityType: EntityType, actionType: ActionType, entityId: number): string => {
    return `${entityType}-${actionType}-${entityId}`;
}

export const buildUuid = (uniqueIdBasedOnEntityAndAction: string): string => {
    return `${uniqueIdBasedOnEntityAndAction}_${uuidv4()}`;
}