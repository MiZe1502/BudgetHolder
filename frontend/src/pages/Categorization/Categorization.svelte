<script lang="typescript">
    import {onMount} from "svelte";
    import {mockData, mockCategories} from "./data";
    import CategoryRow from "./CategoryRow/CategoryRow.svelte";
    import CategoriesContainer
        from "./CategoriesContainer/CategoriesContainer.svelte";
    import {
        categoriesStatus,
        categories,
        simpleCategories,
        simpleCategoriesStatus,
        categoriesTotal,
        findParentCategory, loadCategoriesTree
    } from "../../stores/categories";
    import {LoadingStatus} from "../../stores/utils";

    onMount(async () => {
        await loadCategoriesTree();
    })

</script>

<section>
    <CategoriesContainer categories={$categories} status={$categoriesStatus}/>
</section>
