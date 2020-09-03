<script lang="typescript">
    import {
        middleRange,
        maxRecordsPerPage,
        Delimeters,
        formPagesArray
    } from "./utils";
    import {style} from "./style";
    import {Direction} from "../PaginationArrowElement/utils";

    import {onMount} from 'svelte';

    import PaginationSingleElement
        from "../PaginationSingleElement/PaginationSingleElement.svelte";
    import PaginationDelimeter
        from "../PaginationDelimeter/PaginationDelimeter.svelte";
    import PaginationArrowElement
        from "../PaginationArrowElement/PaginationArrowElement.svelte";

    const changePage = (page: number) => {
        currentPage = page
        pagesInMiddle = formPagesArray(totalPages, currentPage);
        onPageChange(currentPage);
    }

    const nextPageHandler = () => {
        changePage(currentPage + 1);
    }

    const previousPageHandler = () => {
        changePage(currentPage - 1);
    }

    let pagesInMiddle = [];

    $: totalPages = Math.ceil(totalCount / maxRecordsPerPage);

    let currentPage = 1;

    onMount(() => {
        pagesInMiddle = formPagesArray(totalPages, currentPage)
    });

    export let totalCount: number = 0;
    export let onPageChange: () => {};
</script>

<div class="{style.SideMainPadding} {style.Pagination}">
    <div class="{style.FlexHorCenter}">
        <PaginationArrowElement disabled={currentPage === 1}
                                direction={Direction.Left}
                                onClickHandler={previousPageHandler}/>
        <PaginationSingleElement className={style.FirstPaginationElement}
                                 isActive={1 === currentPage} pageNumber={1}
                                 onClick={changePage}/>
        {#each pagesInMiddle as page}
            {#if page === Delimeters.Left || page === Delimeters.Right}
                <PaginationDelimeter/>
            {:else}
                <PaginationSingleElement isActive={page === currentPage}
                                         pageNumber={page}
                                         onClick={changePage}/>
            {/if}
        {/each}
        <PaginationSingleElement isActive={totalPages === currentPage}
                                 pageNumber={totalPages} onClick={changePage}/>
        <PaginationArrowElement disabled={currentPage === totalPages}
                                direction={Direction.Right}
                                onClickHandler={nextPageHandler}/>
    </div>
</div>
