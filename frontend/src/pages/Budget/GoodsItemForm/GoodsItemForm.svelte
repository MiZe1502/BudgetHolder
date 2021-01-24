<script lang="ts">
    import {_} from 'svelte-i18n'
    import {onMount} from "svelte";
    import {style} from "./style";
    import {GoodsData} from "../types";
    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import {isFieldInvalid} from "../../../common/utils/validation";
    import {
        openedPopups,
        updatePopupInnerValidation
    } from "../../../stores/popup";
    import {validationRules} from "./validationRules";
    import InputDropdown
        from "../../../common/components/Inputs/InputDropdown/InputDropdown.svelte";
    import {getSimpleCategoryById} from "../../Categorization/categories";
    import {simpleCategories} from "../../Categorization/categories";
    import TextAreaWithLabel
        from "../../../common/components/Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";

    const validateForm = () => {
        //Dirty hack to change data before on:input
        if (event) {
            goodsItem[event.target.name] = event.target.value;
        }

        for (let rule of validationRules) {
            const isInvalid = rule.validator(goodsItem);
            updatePopupInnerValidation(outerPopupUuid, rule.message, rule.fieldName, isInvalid)
        }
    }

    const onCategorySelect = (selectedId: number) => {
        goodsItem.category = getSimpleCategoryById(selectedId);
        goodsItem.category_id = selectedId;
    }

    onMount(() => {
        validateForm(null);
    })

    $: formErrors = $openedPopups.filter((popup) => popup.uuid === outerPopupUuid).length > 0 && $openedPopups.filter((popup) => popup.uuid === outerPopupUuid)[0].innerValidationErrors;

    export let goodsItem: GoodsData = {
        category: {}
    };
    export let outerPopupUuid: string = "";
</script>

<form class="{style.SideMinorPadding} {style.FlexVert} {style.Form}">
    <InputWithLabel invalid={isFieldInvalid("name", formErrors)}
                    on:input={validateForm} on:change={validateForm}
                    label={$_("common.labels.name")} autofocus={true}
                    type="text" name="name"
                    bind:value={goodsItem.name} required={true}/>
    <InputDropdown
            onSelectHandler={(selectedId) => onCategorySelect(selectedId)}
            bind:value={goodsItem.category.id}
                    name="category"
            label={$_("budget.labels.category")}
            data={$simpleCategories}/>
    <TextAreaWithLabel
            textAreaClass={style.TextArea}
            label={$_("common.labels.comment")}
            name="comment"
            bind:value={goodsItem.comment}/>
</form>

