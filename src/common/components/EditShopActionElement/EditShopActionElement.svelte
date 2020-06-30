<script lang="typescript">
    import { Shop } from "../../../pages/types";
    
    import { EntityType, ActionType } from "../../../stores/utils";

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
        <PopupContainer entityType={EntityType.Shop} actionType={ActionType.Remove} entityId={data.id} title={`Edit ${data.name}` || "Edit"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}>
            Edit Shop Form
        </PopupContainer>
    {/if}
</div>