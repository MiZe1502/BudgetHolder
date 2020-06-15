<script lang="typescript">
    export let onClickHandler = () => {};
    export let title: string = "";
</script>

<style>
    .Button {
        width: auto;
        height: auto;
        background-color: #46b8cc; 
        text-decoration: none;
        border: none;
        outline: none;
        padding: 8px;
        border-radius: 4px;
    }
</style>

<button class="Button Font316Black" onClick={onClickHandler}>
    {title}
</button>