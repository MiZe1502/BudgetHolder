<script lang="typescript">
    import { Shop } from "../../../pages/types";
    import { removeShopFromStore } from "../../../stores/shops";
    import { EntityType, ActionType } from "../../../stores/utils";

    import { Popup, Font732Black, Message, SideMainPadding, FlexHorCenter } from "./style";

    import ButtonIconRemove from "../Buttons/ButtonIconRemove/ButtonIconRemove.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"

    let isPopupOpened = false;

    const onAcceptHandler = () => {
        removeShopFromStore(data.id);
    }

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
        <PopupContainer entityType={EntityType.Shop} actionType={ActionType.Remove} entityId={data.id} onAcceptHandler={onAcceptHandler} onCancelHandler={onCloseHandler} withAcceptButton={true} withCancelButton={true} title={`Remove ${data.name}` || "Remove"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler} popupClass={Popup}>
            <div class="{SideMainPadding} {FlexHorCenter} {Message} {Font732Black}">
                Do You really want to remove {data.name}?
            </div>
        </PopupContainer>
    {/if}
</div>