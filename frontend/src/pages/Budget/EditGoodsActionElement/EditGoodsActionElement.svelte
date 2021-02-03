<script lang="ts">
    import {_} from 'svelte-i18n'
    import {style} from "./style";
    import {GoodsData} from "../types";
    import {onMount} from "svelte";
    import {
        updateGoodsItemInStore,
        addGoodsToStore, addGoodsItemToStore
    } from "../goods";
    import Button
        from "../../../common/components/Buttons/Button/Button.svelte";
    import ButtonIconEdit
        from "../../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {EntityType, ActionType} from "../../../stores/utils";
    import GoodsItemForm from "../GoodsItemForm/GoodsItemForm.svelte";
    import {
        convertSimpleData,
        convertSimpleDataIntoReq
    } from "../../../common/utils/api";

    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
        data = {...initialData}
    }

    const onSaveHandler = async () => {
        const convertedData = {...data, category: convertSimpleDataIntoReq(data.category)}
        if (convertedData.id) {
            await updateGoodsItemInStore(convertedData);
        } else {
            await addGoodsItemToStore(convertedData);
        }
    }

    onMount(() => {
        initialData = {...data};
    })

    let initialData: GoodsData = {};

    export let data: GoodsData = {
        category: {}
    }
    export let withButton: boolean = false;
    export let buttonTitle: string = "";
</script>

<div>
    {#if withButton}
        <Button title={buttonTitle} onClickHandler={onClickHandler}/>
    {:else}
        <ButtonIconEdit onClickHandler={onClickHandler}/>
    {/if}
    {#if isPopupOpened}
    <PopupContainer let:outerPopupUuid={uuid}
                    popupClass={style.Popup} onAcceptHandler={onSaveHandler}
                    onCancelHandler={onCloseHandler} withAcceptButton={true}
                    withCancelButton={$_("common.components.buttons.cancel")}
                    entityType={EntityType.Goods}
                    actionType={ActionType.Edit} entityId={data.id}
                    title={data.id ? `${$_("common.titles.edit")} ${data.name}` || $_("common.titles.edit") : $_("shops.titles.new")}
                    isPopupOpened={isPopupOpened}
                    onCloseHandler={onCloseHandler}
                    acceptButtonTitle={$_("common.components.buttons.save")}>
        <GoodsItemForm outerPopupUuid={uuid} goodsItem={data}/>
    </PopupContainer>
    {/if}
</div>