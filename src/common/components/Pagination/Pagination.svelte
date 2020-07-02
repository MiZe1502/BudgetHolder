<script lang="typescript">
    import { SideMainPadding } from "./style"; 

    import PaginationSingleElement from "../PaginationSingleElement/PaginationSingleElement.svelte";

    const range = (size, startAt = 0) => {
        return [...Array(size).keys()].map(i => i + startAt);
    }

    const changePage = (page: number) => {
        currentPage = page
        console.log(currentPage)
    }

    const maxRecordsPerPage = 10;

    $: totalPages = Math.ceil(totalCount / maxRecordsPerPage);

    let currentPage = 1;

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
        {#each range(totalPages, 1) as pageNumber}
            <PaginationSingleElement isActive={pageNumber === currentPage} pageNumber={pageNumber} onClick={changePage}/>
        {/each}
    </div>
</div>
