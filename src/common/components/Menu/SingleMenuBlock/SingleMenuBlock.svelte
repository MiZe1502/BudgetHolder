<script lang="typescript">
    import { SingleMenuBlock, Font422Black, FlexHorCenter, ActiveBlock } from "./style";

    import { Link } from 'svelte-routing'
    import { currentPage } from "../../../../stores/common";

    export let title = ''
    export let url = ''

    function handleClick (url) {
      currentPage.set(url)
    }
</script>

<Link class="" to={url}>
    <div on:click={() => handleClick(url)} class="{$currentPage === url && ActiveBlock} {SingleMenuBlock} {Font422Black} {FlexHorCenter}">
        {title}
    </div>
</Link>
