<script lang="ts">
    import {onMount} from "svelte";
    import {Shop} from "../types";
    import {isFieldInvalid} from "../../../common/utils/validation";
    import {
        SideMinorPadding,
        TextArea,
        FlexVert,
        Form
    } from "./style";
    import InputWithLabel from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import TextAreaWithLabel
        from "../../../common/components/Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";
    import InputWithClearButton
        from "../../../common/components/Inputs/InputWithClearButton/InputWithClearButton.svelte";
    import InputWithMap from "../../../common/components/Inputs/InputWithMap/InputWithMap.svelte";

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
        console.log("shop", shop)
    })

    $: formErrors = $openedPopups.filter((popup) => popup.uuid === outerPopupUuid).length > 0 && $openedPopups.filter((popup) => popup.uuid === outerPopupUuid)[0].innerValidationErrors;

    export let shop: Partial<Shop> = {};
    export let outerPopupUuid: string = "";
</script>

<style>

</style>

<form class="{SideMinorPadding} {FlexVert} {Form}">
    <InputWithLabel invalid={isFieldInvalid("name", formErrors)}
                          on:input={validateForm} on:change={validateForm}
                          label="Name" autofocus={true}
                          type="text" name="name"
                          bind:value={shop.name} required={true}/>
    <InputWithLabel label="Url" type="text" name="url" bind:value={shop.url}/>
    <InputWithMap label="Address" type="text" name="address"
                  bind:value={shop.address} bind:data={shop}/>
    <TextAreaWithLabel invalid={isFieldInvalid("comment", formErrors)}
                       on:input={validateForm} on:change={validateForm}
                       textAreaClass={TextArea} label="Comment" name="comment"
                       bind:value={shop.comment}/>
</form>