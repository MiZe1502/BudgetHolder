<script lang="ts">
    import {_} from 'svelte-i18n'
    import {onMount} from "svelte";
    import {dateToDMY} from "../../../common/components/ElementsAndBlocks/SimpleDateElement/utils";
    import ButtonIconEdit
        from "../../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {Purchase} from "../types";
    import {LoadingStatus} from "../../../stores/utils";
    import {EntityType, ActionType} from "../../../stores/utils";
    import {updatePurchaseDataInStore} from "../../../stores/purchases";
    import {style} from "./style";
    import PurchaseEditForm from "../PurchaseEditForm/PurchaseEditForm.svelte";

    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
        data = {...initialData}
    }

    const onSaveHandler = () => {
        updatePurchaseDataInStore(data);
    }

    onMount(() => {
        initialData = {...data};
    })

    let initialData: Purchase = {};

    export let data: Purchase = {};
    export let status: LoadingStatus = LoadingStatus.None;
</script>

<div>
    <ButtonIconEdit onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer let:outerPopupUuid={uuid}
                        popupClass={style.Popup} onAcceptHandler={onSaveHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={$_("common.components.buttons.cancel")}
                        entityType={EntityType.Purchase}
                        actionType={ActionType.Edit} entityId={data.id}
                        title={`${$_("budget.titles.edit")}${dateToDMY(data.date)}`}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle={$_("common.components.buttons.save")}>
            <PurchaseEditForm outerPopupUuid={uuid} data={data}/>
        </PopupContainer>
    {/if}
</div>