<script lang="typescript">
    import {onMount} from "svelte";
    import {SideMainPadding} from "./style";
    import {mockData} from "./data";
    import CategoryRow from "./CategoryRow/CategoryRow.svelte";
    import CategoriesContainer
        from "./CategoriesContainer/CategoriesContainer.svelte";

    onMount(() => {

    })

</script>

<section class="{SideMainPadding}">
    <CategoriesContainer categories={mockData}/>
</section>
