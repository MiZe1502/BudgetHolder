<script lang="ts">
    import {_} from 'svelte-i18n'
    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import {onMount} from "svelte";
    import InputDropdown
        from "../../../common/components/Inputs/InputDropdown/InputDropdown.svelte";
    import TextAreaWithLabel
        from "../../../common/components/Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";
    import {Purchase} from "../types";
    import {dateToDMY} from "../../../common/components/ElementsAndBlocks/SimpleDateElement/utils";
    import {getShopById, simpleShops} from "../../../stores/shops";
    import {
        openedPopups,
        updatePopupInnerValidation
    } from "../../../stores/popup";
    import {isFieldInvalid} from "../../../common/utils/validation";
    import {style} from "./style";
    import {validationRulesPurchase} from "../BudgetPurchaseForm/validationRules";

    const validateForm = (event: Event<HTMLInputElement>) => {
        //Dirty hack to change data before on:input
        if (event) {
            data[event.target.name] = event.target.value;
        }

        for (let rule of validationRulesPurchase) {
            const isInvalid = rule.validator(data);
            updatePopupInnerValidation(outerPopupUuid, rule.message, rule.fieldName, isInvalid)
        }
    }

    const onShopSelect = (selectedId: number) => {
        data.shop = {...getShopById(selectedId)};
    }

    onMount(() => {
        validateForm(null);
    })

    $: formErrors = $openedPopups.filter((popup) => popup.uuid === outerPopupUuid).length > 0 && $openedPopups.filter((popup) => popup.uuid === outerPopupUuid)[0].innerValidationErrors;

    export let data: Partial<Purchase> = {};
    export let outerPopupUuid: string = "";
</script>


<form class="{style.SideMinorPadding} {style.FlexVert} {style.Form}">
    <InputWithLabel invalid={isFieldInvalid("name", formErrors)}
                    on:input={validateForm} on:change={validateForm}
                    label={$_("budget.labels.price")} autofocus={true}
                    type="text" name="name"
                    bind:value={data.totalPrice} required={true}/>
    <InputWithLabel
            label={$_("budget.labels.date")} autofocus={true}
            type="date" name="date"
            bind:value={data.date} required={true}/>
    <InputDropdown onSelectHandler={onShopSelect}
                   bind:value={data.shop.id} name="shop"
                   label={$_("budget.labels.shop")}
                   data={$simpleShops}/>
    <TextAreaWithLabel
            textAreaClass="{style.TextArea}"
            label={$_("common.labels.comment")} name="comment"
            bind:value={data.comment}/>
</form>