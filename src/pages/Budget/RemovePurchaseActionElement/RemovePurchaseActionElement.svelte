<script lang="ts">
    import {_} from 'svelte-i18n';
    import {dateToDMY} from "../../../common/components/ElementsAndBlocks/SimpleDateElement/utils";
    import ButtonIconRemove
        from "../../../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {Purchase} from "../types";
    import {removePurchaseFromStore} from "../../../stores/purchases";
    import {EntityType, ActionType} from "../../../stores/utils";
    import {style} from "./style";

    let isPopupOpened = false;

    const onAcceptHandler = () => {
        removePurchaseFromStore(data.id);
    }

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: Partial<Purchase> = {};

</script>

<div>
    <ButtonIconRemove onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer entityType={EntityType.Purchase}
                        actionType={ActionType.Remove} entityId={data.id}
                        onAcceptHandler={onAcceptHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={true}
                        title={`${$_("common.titles.remove")} ${dateToDMY(data.date)}, ${data.totalPrice}` || $_("commmon.titles.remove")}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        popupClass={style.Popup}>
            <div class="{style.SideMainPadding} {style.FlexHorCenter} {style.Message} {style.Font732Black}">
                {`${$_("budget.messages.remove_part1")}${$_("budget.messages.remove_part2")}`}
            </div>
        </PopupContainer>
    {/if}
</div>