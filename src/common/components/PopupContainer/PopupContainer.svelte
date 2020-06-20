<script lang="typescript">

    const countPositionDiff = (curPosition: Position): Position => {
        return {
            x: curPosition.x - startMousePosition.x,
            y: curPosition.y - startMousePosition.y 
        }
    }

    const onMouseDownHandler = (event: MouseEvent) => {
        isMouseDown = true;
        if (!isPopupAlreadyMoved) {
            startMousePosition.x = event.clientX;
            startMousePosition.y = event.clientY;
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

        diffPosition = countPositionDiff({x: event.clientX, y: event.clientY})
    }

    const onMouseUpHandler = (event: MouseEvent) => {
        isMouseMoving = false;
        isMouseDown = false;
    }

    const onMouseOutHandler = (event: MouseEvent) => {
        isMouseMoving = false;
        isMouseDown = false;
    }

    let isMouseDown = false;
    let isMouseMoving = false;
    let isPopupAlreadyMoved = false;

    interface Position {
        x: number;
        y: number;
    }

    let startMousePosition: Position = {x: 0, y: 0};
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

</style>

<div style="top: calc(100px + {diffPosition.y}px); left: calc(50% - 450px + {diffPosition.x}px" class="Popup" on:mousedown={onMouseDownHandler} on:mousemove={onMouseMoveHandler} on:mouseup={onMouseUpHandler} on:mouseout={onMouseOutHandler}>
    Test Popup
</div>