<script lang="typescript">
    import { Link } from 'svelte-routing'
    import { currentPage } from "../../../stores/common";

    export let title = ''
    export let url = ''

    function handleClick (url) {
      currentPage.set(url)
    }
</script>

<style lang="scss">
    .SingleMenuBlock {
        height: 5vh;
        width: auto;
        background-color: #80cf47;
        padding-bottom: 8px;
        padding-top: 8px;
        padding-left: 12px;
        padding-right: 12px;
        min-width: 100px;
        justify-content: center;
    }

    .SingleMenuBlock:active,
    .SingleMenuBlock:hover {
        cursor: pointer;
        background-color: #5b9c35;
    }

    .Active {
        background-color: #5b9c35;
    }
</style>


<Link class="" to={url}>
    <div on:click={() => handleClick(url)} class:Active={$currentPage === url} class="SingleMenuBlock Font422Black PrimaryBackground FlexHor">
        {title}
    </div>
</Link>
