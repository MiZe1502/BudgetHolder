<script lang="typescript">
    import { Shop } from "../../../pages/types";
    import { updateShopInStore } from "../../../stores/shops";

    import {Popup} from "./style";
    import {EntityType, ActionType} from "../../../stores/utils";

    import ButtonIconEdit from "../Buttons/ButtonIconEdit/ButtonIconEdit.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"
    import ShopEditForm from "../ShopEditForm/ShopEditForm.svelte";

    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    const onSaveHandler = () => {
        updateShopInStore(data);
    }

    export let data: Shop = {};
</script>

<div>
    <ButtonIconEdit onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer popupClass={Popup} onAcceptHandler={onSaveHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton="Save"
                        withCancelButton="Cancel" entityType={EntityType.Shop}
                        actionType={ActionType.Remove} entityId={data.id}
                        title={`Edit ${data.name}` || "Edit"}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle="Save">
            <ShopEditForm shop={data}/>
        </PopupContainer>
    {/if}
</div>