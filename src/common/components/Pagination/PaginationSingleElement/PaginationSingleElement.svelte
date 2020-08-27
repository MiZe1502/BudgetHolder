<script lang="typescript">
    import {style} from "./style";

    export let pageNumber: number = 0;
    export let onClick: (page: number) => {};
    export let isActive: boolean = false;
    export let className: string = "";
</script>

<div on:click={() => onClick(pageNumber)}
     class="{style.Font312Black} {style.FlexHorCenter} {style.SingleButton} {isActive && style.ActiveButton} {className}">
    {pageNumber}
</div>