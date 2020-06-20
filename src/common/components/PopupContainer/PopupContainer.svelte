<script lang="typescript">

    const onMouseDownHandler = (event: MouseEvent) => {
        console.log("MouseDown", event)
        isMouseDown = true;
        startMousePosition.x = event.clientX;
        startMousePosition.y = event.clientY;
    }

    const onMouseMoveHandler = (event: MouseEvent) => {
        if (!isMouseDown) {
            return;
        }
        console.log("MouseMove", event)
        if (!isMouseMoving) {
            isMouseMoving = true;
        }
    }

    const onMouseUpHandler = (event: MouseEvent) => {
        console.log("MouseUp", event)
        isMouseMoving = false;
        isMouseDown = false;
    }

    let isMouseDown = false;
    let isMouseMoving = false;

    interface Position {
        x: number;
        y: number;
    }

    let startMousePosition: Position = {x: 0, y: 0};



</script>

<style>
    .Popup {
        left: calc(50% - 450px);
        top: 100px;
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

<div class="Popup" on:mousedown={onMouseDownHandler} on:mousemove={onMouseMoveHandler} on:mouseup={onMouseUpHandler}>
    Test Popup
</div>