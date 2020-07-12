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

    const validateForm = () => {
        console.log(outerPopupUuid)

        if (!shop.name) {
            updatePopupInnerValidation(outerPopupUuid, "Name field should not be empty")
        }


        // let isValid = true;
        //
        // if (!shop.name) {
        //     isValid = false;
        // }
        //
        // console.log(isValid)
        //
        // validateHandler(isValid);
    }

    export let shop: Partial<Shop> = {};
    export let outerPopupUuid: string = "";
</script>

<style>

</style>

<form class="{SideMinorPadding} {FlexVert} {Form}">
    <InputWithClearButton onChange={validateForm} label="Name" autofocus={true}
                          type="text" name="name"
                          bind:value={shop.name} required={true}/>
    <InputWithLabel label="Url" type="text" name="url" bind:value={shop.url}/>
    <InputWithMap label="Address" type="text" name="address"
                  bind:value={shop.address} bind:data={shop}/>
    <TextAreaWithLabel textAreaClass={TextArea} label="Comment" name="comment"
                       bind:value={shop.comment}/>
</form>