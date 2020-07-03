<script lang="typescript">
    import { middleRange, maxRecordsPerPage, Delimeters, formPagesArray } from "./utils";
    import { SideMainPadding, FirstPaginationElement } from "./style"; 
    import { onMount } from 'svelte';

    import PaginationSingleElement from "../PaginationSingleElement/PaginationSingleElement.svelte";
    import PaginationDelimeter from "../PaginationDelimeter/PaginationDelimeter.svelte";
    import ButtonIconRightArrow from "../Buttons/ButtonIconRightArrow/ButtonIconRightArrow.svelte";
    import ButtonIconLeftArrow from "../Buttons/ButtonIconLeftArrow/ButtonIconLeftArrow.svelte";

    const changePage = (page: number) => {
        currentPage = page
        pagesInMiddle = formPagesArray(totalPages, currentPage);
    }

    let pagesInMiddle = [];

    $: totalPages = Math.ceil(totalCount / maxRecordsPerPage);

    let currentPage = 1;

    onMount(() => {
        pagesInMiddle = formPagesArray(totalPages, currentPage)
    });

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
        <ButtonIconLeftArrow />
        <PaginationSingleElement className={FirstPaginationElement} isActive={1 === currentPage} pageNumber={1} onClick={changePage}/>
        {#each pagesInMiddle as page}
            {#if page === Delimeters.Left || page === Delimeters.Right}
                <PaginationDelimeter />
            {:else}
                <PaginationSingleElement isActive={page === currentPage} pageNumber={page} onClick={changePage}/>
            {/if}
        {/each}
        <PaginationSingleElement isActive={totalPages === currentPage} pageNumber={totalPages} onClick={changePage}/>
        <ButtonIconRightArrow />
    </div>
</div>
