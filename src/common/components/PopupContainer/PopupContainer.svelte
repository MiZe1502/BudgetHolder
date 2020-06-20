<script lang="typescript">
    import { Font724Black, FlexHorCenter } from "./style";

    const countPositionDiff = (curPosition: Position, prevPosition: Position): Position => {
        return {
            x: curPosition.x - prevPosition.x,
            y: curPosition.y - prevPosition.y 
        }
    }

    const onMouseDownHandler = (event: MouseEvent) => {
        //console.log("mouseDown", {x: event.clientX, y: event.clientY}, startMousePosition, endMousePosition)
        isMouseDown = true;
        currentStartMousePosition.x = event.clientX;
        currentStartMousePosition.y = event.clientY;
        if (!isPopupAlreadyMoved) {
            initialStartMousePosition.x = event.clientX;
            initialStartMousePosition.y = event.clientY;
        } else {
            if (isEndedDragging()) {
                nextStartMousePosition.x = event.clientX;
                nextStartMousePosition.y = event.clientY;
            }
            //endMousePosition = {x: 0, y: 0};
            // startMousePosition.x = startMousePosition.x + (startMousePosition.x - event.clientX);
            // startMousePosition.y = endMousePosition.y + (startMousePosition.y - event.clientY);
        }
    }

    const onMouseMoveHandler = (event: MouseEvent) => {
        if (!isMouseDown) {
            return;
        }

        if (!isPopupAlreadyMoved) {
            isPopupAlreadyMoved = true;
        }

        if (!isMouseMoving) {
            isMouseMoving = true;
        }

        //console.log("mouseMove", {x: event.clientX, y: event.clientY}, startMousePosition, endMousePosition)


        const curPosition = isEndedDragging() ? {x: endMousePosition.x + (event.clientX - currentStartMousePosition.x), y: endMousePosition.y + (event.clientY - currentStartMousePosition.y) }: {x: event.clientX, y: event.clientY}
        // if (isEndedDragging()) {
        //     endMousePosition = {x: 0, y: 0};
        // }


        diffPosition = countPositionDiff(curPosition, initialStartMousePosition)
    }

    const onMouseUpHandler = (event: MouseEvent) => {
        isMouseMoving = false;
        isMouseDown = false;
        endMousePosition = {x: event.clientX, y: event.clientY};
    }

    const onMouseOutHandler = (event: MouseEvent) => {
        if (isMouseDown) {
            endMousePosition = {x: event.clientX, y: event.clientY};
        }
        isMouseMoving = false;
        isMouseDown = false;
    }

    const isEndedDragging = () => {
        return endMousePosition.x !== 0 && endMousePosition.y !== 0;
    }

    let isMouseDown = false;
    let isMouseMoving = false;
    let isPopupAlreadyMoved = false;

    interface Position {
        x: number;
        y: number;
    }

    let initialStartMousePosition: Position = {x: 0, y: 0};
    let currentStartMousePosition: Position = {x: 0, y: 0};
    let nextStartMousePosition: Position = {x: 0, y: 0};
    let endMousePosition: Position = {x: 0, y: 0};
    let diffPosition: Position = {x: 0, y: 0};

</script>

<style>
    .Popup {
        /* left: calc(50% - 450px); */
        /* top: 100px; */
        position: fixed;
        z-index: 10;
        margin: auto;
        width: 900px;
        height: 600px;
        border-radius: 4px;
        background: white;
        box-shadow: 0 0 10px 2px rgba(0, 0, 0, 0.17);
    }

    .Header {
        width: 100%;
        height: 50px;
        background-color: #faf8f6;
        cursor: grab;
        border-top-right-radius: 4px;
        border-top-left-radius: 4px;
        padding: 16px;
        box-sizing: border-box;
    }

</style>

<div style="top: calc(100px + {diffPosition.y}px); left: calc(50% - 450px + {diffPosition.x}px" class="Popup">
    <div style="cursor: {isMouseDown ? "grabbing" : "grab"}" class="{FlexHorCenter} Header" on:mousedown={onMouseDownHandler} on:mousemove={onMouseMoveHandler} on:mouseup={onMouseUpHandler} on:mouseout={onMouseOutHandler}>
        <span class="{Font724Black}">Test Popup</span>
    </div>
    Test Popup Text
</div>