<script lang="ts">
    import {_} from 'svelte-i18n'
    import {onMount} from "svelte";
    import {Popup} from "./style";
    import {EntityType, ActionType} from "../../../stores/utils";
    import {GoodsDetails} from "../types";
    import ButtonIconShowDetails
        from "../../../common/components/Buttons/ButtonIconShowDetails/ButtonIconShowDetails.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import BudgetPurchaseDetails
        from "../BudgetPurchaseDetails/BudgetPurchaseDetails.svelte";

    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
        data = [...initialData]
    }

    onMount(() => {
        initialData = [...data];
        console.log(initialData)
    })

    let initialData: GoodsDetails = [];

    export let data: GoodsDetails = [];
    export let entityId: number = -1;
</script>

<div>
    <ButtonIconShowDetails onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer let:outerPopupUuid={uuid}
                        popupClass={Popup}
                        entityType={EntityType.Purchase}
                        actionType={ActionType.Details} entityId={entityId}
                        title={$_("common.titles.details")}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}>
            <BudgetPurchaseDetails purchaseId={entityId}/>
        </PopupContainer>
    {/if}
</div>