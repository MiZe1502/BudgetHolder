<script lang="ts">
    import {_} from 'svelte-i18n';
    import {dateToDMY} from "../../../common/components/ElementsAndBlocks/SimpleDateElement/utils";
    import ButtonIconRemove
        from "../../../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {GoodsDetails} from "../types";
    import {removeGoodsItemDetailsFromPurchase} from "../purchases";
    import {EntityType, ActionType} from "../../../stores/utils";
    import {style} from "./style";
    import MessageBlock
        from "../../../common/components/ElementsAndBlocks/MessageBlock/MessageBlock.svelte";

    let isPopupOpened = false;

    const onAcceptHandler = async () => {
        await removeGoodsItemDetailsFromPurchase(data.goods_details_id, data.purchase_id);
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
                        onCloseHandler={onCloseHandler}
                        popupClass={style.Popup}>
            <MessageBlock height={300}
                    messageText={`${$_("goods.messages.remove_part1")}${$_("goods.messages.remove_part2")}`}/>
        </PopupContainer>
    {/if}
</div>