<script lang="ts">
    import {_} from 'svelte-i18n';
    import {dateToDMY} from "../../../common/components/ElementsAndBlocks/SimpleDateElement/utils";
    import ButtonIconRemove
        from "../../../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {GoodsDetails} from "../types";
    import {removeGoodsItemDetailsFromPurchase} from "../../../stores/purchases";
    import {EntityType, ActionType} from "../../../stores/utils";
    import {
        Popup,
        Font732Black,
        Message,
        SideMainPadding,
        FlexHorCenter
    } from "./style";

    let isPopupOpened = false;

    const onAcceptHandler = () => {
        removeGoodsItemDetailsFromPurchase(data.id, data.purchaseId);
    }

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: GoodsDetails = {};

</script>

<div>
    <ButtonIconRemove width={16} height={16} onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer entityType={EntityType.Goods}
                        actionType={ActionType.Remove}
                        entityId={data.id}
                        onAcceptHandler={onAcceptHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={true}
                        title={`${$_("common.titles.remove")} ${data.name}` || $_("commmon.titles.remove")}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler} popupClass={Popup}>
            <div class="{SideMainPadding} {FlexHorCenter} {Message} {Font732Black}">
                {`${$_("goods.messages.remove_part1")}${$_("goods.messages.remove_part2")}`}
            </div>
        </PopupContainer>
    {/if}
</div>