<script lang="typescript">
    import {onMount} from "svelte";
    import CategoriesContainer
        from "./CategoriesContainer/CategoriesContainer.svelte";
    import {
        categoriesStatus,
        categories,
        loadCategoriesTree,
        loadSimpleCategoriesList
    } from "../../stores/categories";

    onMount(async () => {
        await loadCategoriesTree();
        await loadSimpleCategoriesList();
    })

</script>

<section>
    <CategoriesContainer categories={$categories} status={$categoriesStatus}/>
</section>
