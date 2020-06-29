<script lang="typescript">
    import { Shop } from "../../../pages/types";

    import ButtonIconRemove from "../Buttons/ButtonIconRemove/ButtonIconRemove.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"

    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: Shop = {};
</script>

<div>
    <ButtonIconRemove onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer title={`Remove ${data.name}` || "Remove"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}>
            Do You really want to remove {data.name}?
        </PopupContainer>
    {/if}
</div>