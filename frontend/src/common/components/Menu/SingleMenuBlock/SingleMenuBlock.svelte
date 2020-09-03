<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {style} from "./style";

    import {Link} from 'svelte-routing'
    import {currentPage} from "../../../../stores/common";

    export let title = ''
    export let url = ''

    function handleClick(url) {
        currentPage.set(url)
    }
</script>

<Link class="" to={url}>
    <div on:click={() => handleClick(url)}
         class="{$currentPage === url && style.ActiveBlock} {style.SingleMenuBlock} {style.Font422Black} {style.FlexHorCenter}">
        {$_(title)}
    </div>
</Link>
