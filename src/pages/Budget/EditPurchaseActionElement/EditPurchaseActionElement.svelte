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
    import {Popup} from "./style";

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
                        popupClass={Popup} onAcceptHandler={onSaveHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={$_("common.components.buttons.cancel")}
                        entityType={EntityType.Purchase}
                        actionType={ActionType.Remove} entityId={data.id}
                        title={`${$_("budget.titles.edit")}${dateToDMY(data.date)}`}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle={$_("common.components.buttons.save")}>
            EDIT PURCHASE
        </PopupContainer>
    {/if}
</div>