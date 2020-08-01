<script lang="ts">
    import {SideMainPadding, Wrapper, Container} from "./style";
</script>

<section class="{SideMainPadding} {Container}">
    <div class={Wrapper}>
        <slot/>
    </div>
</section>