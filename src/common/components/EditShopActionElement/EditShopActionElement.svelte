<script lang="typescript">
    import { Shop } from "../../../pages/types";

    import ButtonIconEdit from "../Buttons/ButtonIconEdit/ButtonIconEdit.svelte"
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
    <ButtonIconEdit onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer title={`Edit ${data.name}` || "Edit"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}>
            Edit Shop Form
        </PopupContainer>
    {/if}
</div>