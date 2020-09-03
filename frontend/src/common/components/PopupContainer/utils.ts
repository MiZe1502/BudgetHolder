
export const defaultZIndex = 10;
export const defaultPopupTopPosition = 100;
export const popupWidth = 900;

export interface Position {
    x: number;
    y: number;
}

export const countPositionDiff = (curPosition: Position, prevPosition: Position): Position => {
    return {
        x: curPosition.x - prevPosition.x,
        y: curPosition.y - prevPosition.y 
    }
}

export const isNotHeader = (event: MouseEvent) => (<HTMLDivElement>event.target).localName !== "div";

