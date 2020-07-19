<script lang="typescript">
    import {onMount} from "svelte";
    import {SideMainPadding} from "./style";
    import {mockData} from "./data";
    import CategoryRow from "./CategoryRow/CategoryRow.svelte";
    import CategoriesContainer
        from "./CategoriesContainer/CategoriesContainer.svelte";
    import {
        categoriesStatus,
        categories,
        categoriesTotal
    } from "../../stores/categories";
    import {LoadingStatus} from "../../stores/utils";

    onMount(() => {
        categoriesStatus.set(LoadingStatus.Loading)

        setTimeout(() => {
            categoriesStatus.set(LoadingStatus.Finished)
            categories.set(mockData);
            categoriesTotal.set(mockData.length);
        }, 5000)
    })

</script>

<section class="{SideMainPadding}">
    <CategoriesContainer categories={$categories} status={$categoriesStatus}/>
</section>
