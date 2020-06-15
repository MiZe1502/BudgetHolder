<script lang="typescript">
    export let onClickHandler = () => {};
    export let title: string = "";
</script>

<style>
    .Button {
        width: auto;
        height: auto;
        background-color: #80cf47; 
        text-decoration: none;
        border: none;
        outline: none;
        padding: 8px;
        border-radius: 4px;
        transition: background-color 300ms ease;
        box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.3);
    }

    .Button:hover,
    .Button:active {
        background-color: #5b9c35;
        cursor: pointer;
        box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.15);
    }

    .Button:active {
        background-color: #9ef06e;
        box-shadow: none;
    }

    .Button:disabled {
        background-color: #d6d2ce;
        border: none;
        cursor: default;
        box-shadow: none;
        color: whitesmoke;
    }
</style>

<button class="Button Font316Black" onClick={onClickHandler}>
    {title}
</button>