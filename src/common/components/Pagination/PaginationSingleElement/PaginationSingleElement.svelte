<script lang="typescript">
    import {
        FlexHorCenter,
        SingleButton,
        ActiveButton,
        Font312Black
    } from "./style";

    export let pageNumber: number = 0;
    export let onClick: (page: number) => {};
    export let isActive: boolean = false;
    export let className: string = "";
</script>

<div on:click={() => onClick(pageNumber)}
     class="{Font312Black} {FlexHorCenter} {SingleButton} {isActive && ActiveButton} {className}">
    {pageNumber}
</div>