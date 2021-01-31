<script lang="ts">
    import {GoodsDetails} from "../types";
    import {_} from 'svelte-i18n'
    import {LoadingStatus} from "../../../stores/utils";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import ButtonIconEdit
        from "../../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte";
    import {EntityType, ActionType} from "../../../stores/utils";
    import {onMount} from "svelte";
    import {style} from "./style";
    import EditGoodsItemForm
        from "../EditGoodsItemForm/EditGoodsItemForm.svelte";
    import {updatePurchaseGoodsItemInStore} from "../../../stores/purchases";

    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
        data = {...initialData}
    }

    const onSaveHandler = async () => {
        await updatePurchaseGoodsItemInStore(data);
    }

    onMount(() => {
        initialData = {...data};
    })

    let initialData: GoodsDetails = {};

    export let data: GoodsDetails = {};
    export let status: LoadingStatus = LoadingStatus.None;
</script>

<div>
    <ButtonIconEdit width={16} height={16} onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer let:outerPopupUuid={uuid}
                        popupClass={style.Popup} onAcceptHandler={onSaveHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={$_("common.components.buttons.cancel")}
                        entityType={EntityType.GoodsDetails}
                        actionType={ActionType.Edit} entityId={data.id}
                        title={`${$_("goods.titles.edit")}`}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle={$_("common.components.buttons.save")}>
            <EditGoodsItemForm outerPopupUuid={uuid} data={data}/>
        </PopupContainer>
    {/if}
</div>