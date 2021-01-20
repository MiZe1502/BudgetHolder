<script lang="ts">
    import {_} from 'svelte-i18n';
    import {style} from "./style";
    import ButtonIconRemove
        from "../../../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte";
    import {GoodsData} from "../types";
    import {removeGoodsFromStore} from "../../../stores/goods";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {EntityType, ActionType} from "../../../stores/utils";
    import MessageBlock
        from "../../../common/components/ElementsAndBlocks/MessageBlock/MessageBlock.svelte";

    let isPopupOpened = false;

    const onAcceptHandler = async () => {
        await removeGoodsFromStore(data.id);
    }

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: GoodsData = {};
</script>


<div>
    <ButtonIconRemove onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer entityType={EntityType.Goods}
                        actionType={ActionType.Remove} entityId={data.id}
                        onAcceptHandler={onAcceptHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={true}
                        title={`${$_("common.titles.remove")} ${data.name}` || $_("common.titles.remove")}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        popupClass={style.Popup}>
            <MessageBlock
                    messageText={`${$_("goods.messages.remove_good_item_part1")}${data.name}${$_("goods.messages.remove_good_item_part2")}`} />
        </PopupContainer>
    {/if}
</div>