<script lang="ts">
    import {Shop} from "../../../pages/types";
    import Button from "../Buttons/Button/Button.svelte";
    import {
        Font312DarkGray,
        Font316Black,
        SideMinorPadding,
        TextArea,
        FlexVert,
        Form
    } from "./style";
    import InputWithLabel from "../Inputs/InputWithLabel/InputWithLabel.svelte";
    import TextAreaWithLabel
        from "../Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";

    export let shop: Partial<Shop> = {};
</script>

<style>

</style>

<form class="{SideMinorPadding} {FlexVert} {Form}">
    <InputWithLabel label="Name" autofocus={true} type="text" name="name"
                    bind:value={shop.name}/>
    <InputWithLabel label="Url" type="text" name="url" bind:value={shop.url}/>
    <InputWithLabel label="Address" type="text" name="address"
                    bind:value={shop.address}/>
    <TextAreaWithLabel textAreaClass={TextArea} label="Comment" name="comment"
                       bind:value={shop.comment}/>
</form>