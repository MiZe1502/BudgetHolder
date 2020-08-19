<script lang="ts">
    import {_} from 'svelte-i18n'
    import {
        SideMinorPadding,
        FlexVert,
        Form,
        TextArea
    } from "./style";
    import {SuggestionItem} from "../../../common/components/Inputs/InputWithSuggestion/utils";
    import InputWithSuggestion
        from "../../../common/components/Inputs/InputWithSuggestion/InputWithSuggestion.svelte";
    import {GoodsDetails} from "../types";
    import {onMount} from "svelte";
    import {
        openedPopups,
        updatePopupInnerValidation
    } from "../../../stores/popup";
    import {validationRulesGoods} from "../BudgetPurchaseForm/validationRules";
    import {goods, goodsForSuggestions} from "../../../stores/goods";
    import {simpleCategories} from "../../../stores/categories";
    import InputDropdown
        from "../../../common/components/Inputs/InputDropdown/InputDropdown.svelte";
    import {getSimpleCategoryById} from "../../../stores/categories";
    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import TextAreaWithLabel
        from "../../../common/components/Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";

    const validateForm = (event: Event<HTMLInputElement>) => {
        //Dirty hack to change data before on:input
        if (event) {
            data[event.target.name] = event.target.value;
        }

        for (let rule of validationRulesGoods) {
            const isInvalid = rule.validator(data);
            updatePopupInnerValidation(outerPopupUuid, rule.message, rule.fieldName, isInvalid)
        }
    }

    onMount(() => {
        validateForm(null);
    })

    const onGoodsItemSelect = (selectedItem: SuggestionItem) => {
        const goodsItemData = $goods.find((item) => item.id === selectedItem.id);
        data.name = goodsItemData.name;
        data.category = goodsItemData.category;
        data.comment = goodsItemData.comment;
    }

    const onCategorySelect = (selectedId: number) => {
        data.category = getSimpleCategoryById(selectedId);
    }

    $: formErrors = $openedPopups.filter((popup) => popup.uuid === outerPopupUuid).length > 0 && $openedPopups.filter((popup) => popup.uuid === outerPopupUuid)[0].innerValidationErrors;

    export let data: Partial<GoodsDetails> = {};
    export let outerPopupUuid: string = "";
</script>

<form class="{SideMinorPadding} {FlexVert} {Form}">
    <InputWithSuggestion
            onSelectHandler={(selectedItem) => onGoodsItemSelect(selectedItem)}
            suggestionsList={$goodsForSuggestions}
                    label={$_("common.labels.name")}
                    type="text" name="name"
            bind:value={data.name}/>
    <InputDropdown
            onSelectHandler={(selectedId) => onCategorySelect(selectedId)}
            bind:value={data.category.id}
                    name="category"
            label={$_("budget.labels.category")}
                    data={$simpleCategories}/>
    <InputWithLabel
            label={$_("budget.labels.amount")}
            type="number" name="amount"
            bind:value={data.amount}/>
    <InputWithLabel
            label={$_("budget.labels.single_price")}
            type="number" name="price"
            bind:value={data.price}/>
    <TextAreaWithLabel
            textAreaClass={TextArea}
            label={$_("common.labels.comment")}
            name="comment"
            bind:value={data.comment}/>
</form>