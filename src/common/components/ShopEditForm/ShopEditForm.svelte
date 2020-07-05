<script lang="ts">
    import {Shop} from "../../../pages/types";
    import Button from "../Buttons/Button/Button.svelte";

    const shop: Partial<Shop> = {
        name: "",
        url: "",
        address: "",
        comment: "",
    }

    const onSubmit = (e: MouseEvent) => {
        e.preventDefault();
        console.log(shop)
    }

</script>

<style>
    .Form {
        display: flex;
        flex-direction: column;
        padding-top: 8px;
        padding-bottom: 8px;
    }

    .Input {
        height: 30px;
        border-radius: 4px;
        padding: 4px 8px;
        border: 1px solid #80cf47;
        margin-bottom: 8px;
        caret-color: #d6d2ce;
    }

    .Text {
        width: 100%;
        height: 50px;
    }
</style>

<form class="Form">
    <label for="name">Name</label>
    <input class="Input" type="text" name="name" bind:value={shop.name} autofocus/>

    <label for="url">Url</label>
    <input class="Input" type="text" name="url" bind:value={shop.url} />

    <label for="address">Address</label>
    <input class="Input" type="text" name="address" bind:value={shop.address} />

    <textarea class="Text" name="comment" bind:value={shop.comment} />

    <Button title="OK" onClickHandler={onSubmit} />
</form>