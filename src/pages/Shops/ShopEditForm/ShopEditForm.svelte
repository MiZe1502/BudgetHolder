<script lang="ts">
    import {_} from 'svelte-i18n'
    import {onMount} from "svelte";
    import {Shop} from "../types";
    import {isFieldInvalid} from "../../../common/utils/validation";
    import {style} from "./style";
    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import TextAreaWithLabel
        from "../../../common/components/Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";
    import InputWithClearButton
        from "../../../common/components/Inputs/InputWithClearButton/InputWithClearButton.svelte";
    import InputWithMap
        from "../../../common/components/Inputs/InputWithMap/InputWithMap.svelte";

    import {
        openedPopups,
        updatePopupInnerValidation
    } from "../../../stores/popup";
    import {validationRules} from "./validationRules";

    const validateForm = (event: Event<HTMLInputElement>) => {
        //Dirty hack to change data before on:input
        if (event) {
            shop[event.target.name] = event.target.value;
        }

        for (let rule of validationRules) {
            const isInvalid = rule.validator(shop);
            updatePopupInnerValidation(outerPopupUuid, rule.message, rule.fieldName, isInvalid)
        }
    }

    onMount(() => {
        validateForm(null);
    })

    $: formErrors = $openedPopups.filter((popup) => popup.uuid === outerPopupUuid).length > 0 && $openedPopups.filter((popup) => popup.uuid === outerPopupUuid)[0].innerValidationErrors;

    export let shop: Partial<Shop> = {};
    export let outerPopupUuid: string = "";
</script>

<form class="{style.SideMinorPadding} {style.FlexVert} {style.Form}">
    <InputWithLabel invalid={isFieldInvalid("name", formErrors)}
                    on:input={validateForm} on:change={validateForm}
                    label={$_("common.labels.name")} autofocus={true}
                    type="text" name="name"
                    bind:value={shop.name} required={true}/>
    <InputWithLabel label={$_("common.labels.url")} type="text" name="url"
                    bind:value={shop.url}/>
    <InputWithMap label={$_("common.labels.address")} type="text" name="address"
                  bind:value={shop.address} bind:data={shop}/>
    <TextAreaWithLabel invalid={isFieldInvalid("comment", formErrors)}
                       on:input={validateForm} on:change={validateForm}
                       textAreaClass={style.TextArea}
                       label={$_("common.labels.comment")} name="comment"
                       bind:value={shop.comment}/>
</form>