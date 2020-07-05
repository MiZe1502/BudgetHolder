<script lang="ts">
    import {Shop} from "../../../pages/types";
    import Button from "../Buttons/Button/Button.svelte";
    import {Font312DarkGray, Font316Black, SideMinorPadding} from "./style";
    import InputWithLabel from "../Inputs/InputWithLabel/InputWithLabel.svelte";

    export let shop: Partial<Shop> = {};
</script>

<style>
    .Form {
        display: flex;
        flex-direction: column;
        padding-top: 8px;
        padding-bottom: 8px;
    }

    .Text {
        width: 100%;
        height: 50px;
    }
</style>

<form class="{SideMinorPadding} Form">
    <InputWithLabel label="Name" autofocus={true} type="text" name="name" bind:value={shop.name}/>
    <InputWithLabel label="Url" type="text" name="url" bind:value={shop.url}/>
    <InputWithLabel label="Address" type="text" name="address" bind:value={shop.address}/>

    <textarea class="Text" name="comment" bind:value={shop.comment} />
</form>