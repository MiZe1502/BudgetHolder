<script lang="typescript">
    import {onMount} from "svelte";
    import {SideMainPadding} from "./style";
    import {mockData, mockCategories} from "./data";
    import CategoryRow from "./CategoryRow/CategoryRow.svelte";
    import CategoriesContainer
        from "./CategoriesContainer/CategoriesContainer.svelte";
    import {
        categoriesStatus,
        categories,
        simpleCategories,
        simpleCategoriesStatus,
        categoriesTotal
    } from "../../stores/categories";
    import {LoadingStatus} from "../../stores/utils";

    onMount(() => {
        categoriesStatus.set(LoadingStatus.Loading)
        simpleCategoriesStatus.set(LoadingStatus.Loading)

        setTimeout(() => {
            categoriesStatus.set(LoadingStatus.Finished)
            simpleCategoriesStatus.set(LoadingStatus.Finished)

            categories.set(mockData);
            simpleCategories.set(mockCategories);

            categoriesTotal.set(mockData.length);
        }, 5000)
    })

</script>

<section>
    <CategoriesContainer categories={$categories} status={$categoriesStatus}/>
</section>
