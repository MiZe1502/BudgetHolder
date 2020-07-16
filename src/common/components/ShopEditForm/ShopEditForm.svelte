<script lang="ts">
    import {Shop} from "../../../pages/types";
    import {
        SideMinorPadding,
        TextArea,
        FlexVert,
        Form
    } from "./style";
    import InputWithLabel from "../Inputs/InputWithLabel/InputWithLabel.svelte";
    import TextAreaWithLabel
        from "../Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";
    import InputWithClearButton
        from "../Inputs/InputWithClearButton/InputWithClearButton.svelte";
    import InputWithMap from "../Inputs/InputWithMap/InputWithMap.svelte";

    import {
        updatePopupInnerValidation
    } from "../../../stores/popup";
    import {validationRules} from "./validationRules";

    const validateForm = (event: Event<HTMLInputElement>) => {
        //Dirty hack to change data before on:input
        shop[event.target.name] = event.target.value;

        for (let rule of validationRules) {
           const isInvalid = rule.validator(shop);
           updatePopupInnerValidation(outerPopupUuid, rule.message, rule.fieldName, isInvalid)
        }
    }

    export let shop: Partial<Shop> = {};
    export let outerPopupUuid: string = "";
</script>

<style>

</style>

<form class="{SideMinorPadding} {FlexVert} {Form}">
    <InputWithClearButton on:input={validateForm} on:change={validateForm}
                          label="Name" autofocus={true}
                          type="text" name="name"
                          bind:value={shop.name} required={true}/>
    <InputWithLabel label="Url" type="text" name="url" bind:value={shop.url}/>
    <InputWithMap label="Address" type="text" name="address"
                  bind:value={shop.address} bind:data={shop}/>
    <TextAreaWithLabel on:input={validateForm} on:change={validateForm}
                       textAreaClass={TextArea} label="Comment" name="comment"
                       bind:value={shop.comment}/>
</form>