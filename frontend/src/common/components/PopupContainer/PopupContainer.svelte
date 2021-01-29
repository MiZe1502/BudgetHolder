<script lang="typescript">
    import {_} from 'svelte-i18n'
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
    import {style} from "./style";

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

    const onPopupAcceptHandler = async (event: MouseEvent) => {
        event.preventDefault();

        await onAcceptHandler();
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
    export let acceptButtonTitle: string = $_("common.components.buttons.ok");
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
     class="{style.Popup} {popupClass}">
    <div style="cursor: {isMouseDown ? "grabbing" : "grab"}" class="{style.UtilBlock} {style.FlexHorCenter} {style.Header}" on:mousedown={onHeaderMouseDownHandler} on:mousemove={onHeaderMouseMoveHandler} on:mouseup={onHeaderMouseUpHandler} on:mouseout={onHeaderMouseOutHandler}>
    <span class="{style.Font724Black} {style.Overflowed} {style.HeaderText}">{title}</span>
    <ButtonIconClose onClickHandler={onPopupCloseHandler}/>
</div>
<slot outerPopupUuid={uuid}/>
{#if withAcceptButton || withCancelButton}
    <div class="{style.UtilBlock} {style.FlexHorCenter} {style.Footer} {innerComponentErrors && style.FooterWithErrors}">
        {#if innerComponentErrors}
            <ul class="{style.FlexVert} {style.Font312Red} {style.ErrorsBlock}">
                {#each innerComponentErrors as errorField}
                    {#each errorField.errors as error}
                        <li>{error}</li>
                    {/each}
                {/each}
            </ul>
        {/if}
        <div class="{style.FlexHorCenter}">
            {#if withAcceptButton}
            <Button disabled={innerComponentErrors.length > 0}
                    buttonClass={style.PopupButton}
                            title={acceptButtonTitle}
                            onClickHandler={onPopupAcceptHandler}/>
            {/if}
            {#if withCancelButton}
                <Button secondary={true} buttonClass={style.PopupButton}
                        title={$_("common.components.buttons.cancel")}
                        onClickHandler={onPopupCancelHandler}/>
            {/if}
        </div>
    </div>
{/if}
        </div>