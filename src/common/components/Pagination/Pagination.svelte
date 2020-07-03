<script lang="typescript">
    import { SideMainPadding } from "./style"; 
    import { onMount } from 'svelte';

    import PaginationSingleElement from "../PaginationSingleElement/PaginationSingleElement.svelte";

    const range = (size, startAt = 0) => {
        return [...Array(size).keys()].map(i => i + startAt);
    }

    const formPagesArray = () => {
        pagesInMiddle = [];
        if (totalPages > 2) {
            const anchorPage = Math.max(2, Math.min(currentPage - 2, totalPages - 5));
            if (anchorPage > 2) {
                pagesInMiddle.push("L");
            }
            for (let i = 0; i <= 4; i += 1) {
                const page = anchorPage + i;
                if (page < totalPages) {
                    pagesInMiddle.push(page);
                }
            }
            if (anchorPage + 5 < totalPages) {
                pagesInMiddle.push("R");
            }
        }
    }

    const changePage = (page: number) => {
        currentPage = page
        formPagesArray();
        console.log(currentPage)
    }

    const maxRecordsPerPage = 10;
    let pagesInMiddle = [];

    $: totalPages = Math.ceil(totalCount / maxRecordsPerPage);

    let currentPage = 1;

    onMount(() => formPagesArray());

    export let totalCount: number = 0;
</script>

<style>
    .Pagination {
        display: flex;
        flex-direction: row;
        align-items: center;
    }

    .Wrapper {
        width: 100%;
    }

</style>

<div class="Wrapper {SideMainPadding}">
    <div class="Pagination">
        <PaginationSingleElement isActive={1 === currentPage} pageNumber={1} onClick={changePage}/>
        {#each pagesInMiddle as page}
            {#if page === "L"}
                L
            {:else}
                {#if page === "R"}
                    R
                {:else}
                    <PaginationSingleElement isActive={page === currentPage} pageNumber={page} onClick={changePage}/>
                {/if}
            {/if}
        {/each}
        <PaginationSingleElement isActive={totalPages === currentPage} pageNumber={totalPages} onClick={changePage}/>
    </div>
</div>
