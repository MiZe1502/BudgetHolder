<script lang="ts">
    import {_} from 'svelte-i18n'
    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import {Category} from "../types";
    import {isFieldInvalid} from "../../../common/utils/validation";
    import {style} from "./style";
    import {
        openedPopups,
        updatePopupInnerValidation
    } from "../../../stores/popup";
    import {simpleCategories} from "../../../stores/categories";
    import {onMount} from "svelte";
    import {validationRules} from "./validationRules";
    import InputDropdown
        from "../../../common/components/Inputs/InputDropdown/InputDropdown.svelte";

    const validateForm = (event: Event<HTMLInputElement>) => {
        //Dirty hack to change data before on:input
        if (event) {
            data[event.target.name] = event.target.value;
        }

        for (let rule of validationRules) {
            const isInvalid = rule.validator(data);
            updatePopupInnerValidation(outerPopupUuid, rule.message, rule.fieldName, isInvalid)
        }
    }

    onMount(() => {
        validateForm(null);
    })

    const onSelectParentIdHandler = (selectedId: number) => {
        data.parent_id = selectedId;
        console.log("onSelectParentIdHandler", data)
    }


    $: formErrors = $openedPopups.filter((popup) => popup.uuid === outerPopupUuid).length > 0 && $openedPopups.filter((popup) => popup.uuid === outerPopupUuid)[0].innerValidationErrors;

    export let data: Partial<Category> = {};
    export let outerPopupUuid: string = "";
</script>
<form class="{style.SideMinorPadding} {style.FlexVert} {style.Form}">
    <InputWithLabel invalid={isFieldInvalid("name", formErrors)}
                    on:input={validateForm} on:change={validateForm}
                    label={$_("common.labels.name")} autofocus={true}
                    type="text" name="name"
                    bind:value={data.name} required={true}/>

    <InputDropdown onSelectHandler={onSelectParentIdHandler}
                   bind:value={data.parent_id} name="category" label="Category"
                   data={$simpleCategories}/>
</form>