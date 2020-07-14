<script lang="typescript">
    import {onMount} from 'svelte';
    import {fade} from 'svelte/transition';
    import {get} from 'svelte/store';

    import ButtonIconClose
        from "../Buttons/ButtonIconClose/ButtonIconClose.svelte";
    import Button from "../Buttons/Button/Button.svelte";

    import {
        defaultZIndex,
        popupWidth,
        defaultPopupTopPosition,
        isNotHeader,
        countPositionDiff,
        Position
    } from "./utils";
    import {
        PopupButton,
        UtilBlock,
        Footer,
        Font724Black,
        FlexHorCenter,
        Popup,
        Header,
        Overflowed,
        HeaderText,
        FlexVert,
        Font312Red,
        FooterWithErrors,
        ErrorsBlock
    } from "./style";

    import {
        addPopupToState,
        movePopupToTheTop,
        openedPopups,
        PopupState,
        removePopupFromStore
    } from "../../../stores/popup";
    import {
        buildUniqueIdBasedOnEntityAndAction,
        buildUuid,
        EntityType,
        ActionType
    } from "../../../stores/utils";

    const onPopupMouseDownHandler = (event: MouseEvent) => {
        curPopupState = movePopupToTheTop(curPopupState.uuid);
    }

    const onHeaderMouseDownHandler = (event: MouseEvent) => {
        //TODO: Too dirty. Think how to fix bug when this handler executes on click by svg in close button
        // Svg exactly, not button, not path. Just svg.
        if (isNotHeader(event)) {
            return;
        }

        isMouseDown = true;

        currentPopupPosition = {
            left: (event.target as HTMLDivElement).parentNode.offsetLeft,
            top: (event.target as HTMLDivElement).parentNode.offsetTop
        };

        previousMousePosition = {
            x: event.clientX,
            y: event.clientY,
        }
    }

    const onHeaderMouseMoveHandler = (event: MouseEvent) => {
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

        currentPopupPosition = {
            left: currentPopupPosition.left + diffPosition.x,
            top: currentPopupPosition.top + diffPosition.y
        }

        previousMousePosition = currentMousePosition;
    }

    const onHeaderMouseUpHandler = (event: MouseEvent) => {
        isMouseMoving = false;
        isMouseDown = false;
    }

    const onHeaderMouseOutHandler = (event: MouseEvent) => {
        isMouseMoving = false;
        isMouseDown = false;
    }

    let isMouseDown = false;
    let isMouseMoving = false;

    let currentPopupPosition = {
        top: defaultPopupTopPosition,
        left: window.innerWidth / 2 - popupWidth / 2
    };
    let currentMousePosition: Position = {x: 0, y: 0};
    let previousMousePosition: Position = {x: 0, y: 0};

    let diffPosition: Position = {x: 0, y: 0};

    let curPopupState: PopupState = {
        zIndex: defaultZIndex,
        uuid: "",
    }

    onMount(() => {
        //   const uuid = buildUuid(buildUniqueIdBasedOnEntityAndAction(entityType, actionType, entityId));
        curPopupState = addPopupToState(uuid);
    });

    const onPopupAcceptHandler = (event: MouseEvent) => {
        event.preventDefault();

        onAcceptHandler();
        onCloseHandler();
        removePopupFromStore(curPopupState.uuid);
    }

    const onPopupCloseHandler = (event: MouseEvent) => {
        event.preventDefault();

        onCloseHandler();
        removePopupFromStore(curPopupState.uuid);
    }

    const onPopupCancelHandler = (event: MouseEvent) => {
        event.preventDefault();

        onCancelHandler();
        onCloseHandler();
        removePopupFromStore(curPopupState.uuid);
    }

    export let onCloseHandler = () => {
    };
    export let title: string = "";
    export let popupClass: string = "";
    export let withAcceptButton: boolean = false;
    export let acceptButtonTitle: string = "OK";
    export let withCancelButton: boolean = false;
    export let uniqueId: string = "";
    export let entityType: EntityType = "";
    export let actionType: ActionType = "";
    export let entityId: number;

    $: uuid = buildUuid(buildUniqueIdBasedOnEntityAndAction(entityType, actionType, entityId));
    $: innerComponentErrors = $openedPopups.filter((popup) => popup.uuid === uuid).length > 0 && $openedPopups.filter((popup) => popup.uuid === uuid)[0].innerValidationErrors;

    export let onCancelHandler = () => {
    };
    export let onAcceptHandler = () => {
    };
</script>

<div in:fade="{{duration: 200}}" out:fade="{{duration: 200}}"
     on:mousedown={onPopupMouseDownHandler}
     style="z-index: {curPopupState.zIndex};top: {currentPopupPosition.top}px; left: {currentPopupPosition.left}px"
     class="{Popup} {popupClass}">
    <div style="cursor: {isMouseDown ? "grabbing" : "grab"}" class="{UtilBlock} {FlexHorCenter} {Header}" on:mousedown={onHeaderMouseDownHandler} on:mousemove={onHeaderMouseMoveHandler} on:mouseup={onHeaderMouseUpHandler} on:mouseout={onHeaderMouseOutHandler}>




    <span class="{Font724Black} {Overflowed} {HeaderText}">{title}</span>
    <ButtonIconClose onClickHandler={onPopupCloseHandler}/>
</div>
<slot outerPopupUuid={uuid}/>
{#if withAcceptButton || withCancelButton}
    <div class="{UtilBlock} {FlexHorCenter} {Footer} {innerComponentErrors && FooterWithErrors}">
        {#if innerComponentErrors}
            <ul class="{FlexVert} {Font312Red} {ErrorsBlock}">
                {#each innerComponentErrors as errorField}
                    {#each errorField.errors as error}
                        <li>{error}</li>
                    {/each}
                {/each}
            </ul>
        {/if}
        <div class="{FlexHorCenter}">
            {#if withAcceptButton}
                <Button buttonClass={PopupButton}
                        title={acceptButtonTitle}
                        onClickHandler={onPopupAcceptHandler}/>
            {/if}
            {#if withCancelButton}
                <Button secondary={true} buttonClass={PopupButton} title={"Cancel"}
                        onClickHandler={onPopupCancelHandler}/>
            {/if}
        </div>
    </div>
{/if}
        </div>