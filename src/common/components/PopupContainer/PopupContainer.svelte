<script lang="typescript">
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';

    import ButtonIconClose from "../Buttons/ButtonIconClose/ButtonIconClose.svelte";

    import { Font724Black, FlexHorCenter } from "./style";

    import { addPopupToState, movePopupToTheTop, openedPopups, PopupState, removePopupFromStore } from "../../../stores/popup";

    const countPositionDiff = (curPosition: Position, prevPosition: Position): Position => {
        return {
            x: curPosition.x - prevPosition.x,
            y: curPosition.y - prevPosition.y 
        }
    }

    const onPopupMouseDownHandler = (event: MouseEvent) => {
        curPopupState = movePopupToTheTop(curPopupState.uuid);
    }

    const onMouseDownHandler = (event: MouseEvent) => {
        //TODO: Too dirty. Think how to fix bug when this handler executes on click by svg in close button
        // Svg exactly, not button, not path. Just svg.
        if (event.target.localName !== "div") {
            return;
        }

        isMouseDown = true;

        currentPopupPosition = {
            left: event.target.parentNode.offsetLeft, 
            top: event.target.parentNode.offsetTop
        };

        previousMousePosition = {
            x: event.clientX,
            y: event.clientY,
        }
    }

    const onMouseMoveHandler = (event: MouseEvent) => {
        if (!isMouseDown) {
            return;
        }

        if (!isMouseMoving) {
            isMouseMoving = true;
        }

        currentMousePosition = {
            x: event.clientX, 
            y: event.clientY
        };

        diffPosition = countPositionDiff(currentMousePosition, previousMousePosition);
        previousMousePosition = currentMousePosition;

        currentPopupPosition.left = currentPopupPosition.left + diffPosition.x;
        currentPopupPosition.top = currentPopupPosition.top + diffPosition.y;
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

    interface Position {
        x: number;
        y: number;
    }


    let currentPopupPosition = {top: 100, left: window.innerWidth / 2 - 450};
    let currentMousePosition: Position = {x: 0, y: 0};
    let previousMousePosition: Position = {x: 0, y: 0};

    let diffPosition: Position = {x: 0, y: 0};

    let curPopupState: PopupState = {
        zIndex: 10,
        uuid: "",
    }

    onMount(() => {
        curPopupState = addPopupToState();
    });

    export let onCloseHandler = () => {};

    const onPopupCloseHandler = (event: MouseEvent) => {
        onCloseHandler();
        removePopupFromStore(curPopupState.uuid);
    }
</script>

<style>
    .Popup {
        /* left: calc(50% - 450px); */
        /* top: 100px; */
        position: fixed;
        /* z-index: 10; */
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
        justify-content: space-between;
    }

</style>

<div on:mousedown={onPopupMouseDownHandler} style="z-index: {curPopupState.zIndex};top: {currentPopupPosition.top}px; left: {currentPopupPosition.left}px" class="Popup">
    <div style="cursor: {isMouseDown ? "grabbing" : "grab"}" class="{FlexHorCenter} Header" on:mousedown={onMouseDownHandler} on:mousemove={onMouseMoveHandler} on:mouseup={onMouseUpHandler} on:mouseout={onMouseOutHandler}>
        <span class="{Font724Black}">Test Popup</span>
        <ButtonIconClose onClickHandler={onPopupCloseHandler}/>
    </div>
    Test Popup Text
</div>