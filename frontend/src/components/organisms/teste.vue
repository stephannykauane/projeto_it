<template>
    <div id="app">
        <section class="pen-description">
            <h1>Sign up form <br>with multiple steps in VueJS</h1>
        </section>
        <section class="register" v-if="!hasSeenCongrats">
            <div class="register-icon">
                <img class="register-icon-item" src="https://vuejs.org/images/logo.png" alt="vue logo">

            </div>

            <h2 class="register-title">Sign up for a new account</h2>

            <div class="register-stepper">
                <div class="step" :class="{ 'step-active': step === 1, 'step-done': step > 1 }"><span
                        class="step-number">1</span></div>
                <div class="step" :class="{ 'step-active': step === 2, 'step-done': step > 2 }"><span
                        class="step-number">2</span></div>
                <div class="step" :class="{ 'step-active': step === 3, 'step-done': step > 3 }"><span
                        class="step-number">3</span></div>
            </div>

            <transition name="slide-fade">
                <section v-show="step === 1">
                    <form class="form" method="post" action="#" @submit.prevent="next">
                        <div class="form-group">
                            <input id="loyaltyCivilityMr" type="radio" value="1" v-model="customer.gender">
                            <label class="input-label" for="loyaltyCivilityMr">Mr.</label>

                            <input id="loyaltyCivilityMrs" type="radio" value="2" v-model="customer.gender">
                            <label class="input-label" for="loyaltyCivilityMrs">Mrs.</label>
                        </div>

                        <div class="form-group">
                            <input type="text" v-model="customer.firstName" autocomplete='customer.firstName'
                                placeholder="First name" required />
                            <input type="text" v-model="customer.lastName" autocomplete='customer.lastName'
                                placeholder="Last name" required />
                        </div>

                        <div class="form-group">
                            <input type="tel" v-model="customer.phoneNumber" autocomplete="customer.phoneNumber"
                                placeholder="Phone number" minlength="9" maxlength="10"
                                pattern="^(?:0|\(?\+33\)?\s?|0033\s?)[1-79](?:[\.\-\s]?\d\d){4}$" required />
                        </div>
                        <div class="cta" data-style="see-through" data-alignment="right" data-text-color="custom">
                            <p class="cta-color">
                                <span class="link_wrap">
                                    <input type="submit" value="Next" class="link_text" />
                                    <span class="arrow-next"></span>
                                </span>
                            </p>
                        </div>
                    </form>
                </section>
            </transition>
            <transition name="slide-fade">
                <section v-show="step === 2">
                    <form class="form" method="post" action="#" @submit.prevent="next">
                        <div class="form-group">
                            <input type="text" v-model="customer.address" autocomplete='customer.address'
                                placeholder="Address" required />
                        </div>
                        <div class="form-group">
                            <input type="text" v-model="customer.zipCode" autocomplete='customer.zipCode'
                                placeholder="Zip Code" minlength="5" maxlength="5" required />
                            <input type="text" v-model="customer.city" autocomplete='customer.city' placeholder="City"
                                required />
                        </div>

                        <div class="form-group cta-step">
                            <div class="cta prev">
                                <p class="cta-color">
                                    <span class="link_wrap">
                                        <a class="link_text" href="#" @click.prevent="prev()"><span
                                                class="arrow-prev"></span>Previous
                                        </a>
                                    </span>
                                </p>
                            </div>
                            <div class="cta">
                                <p class="cta-color">
                                    <span class="text"></span>
                                    <span class="link_wrap">
                                        <input type="submit" value="Next" class="link_text" />
                                        <span class="arrow-next"></span>
                                    </span>
                                </p>
                            </div>
                        </div>
                    </form>
                </section>
            </transition>
            <transition name="slide-fade">
                <section v-show="step === 3">
                    <form class="form" action="#" @submit.prevent="customerRegister">
                        <div class="form-group">
                            <input type="email" v-model="customer.eMail" autocomplete="customer.eMail"
                                placeholder="Email" required />
                        </div>
                        <div class="form-group">
                            <input type="date" v-model="customer.birthDay" placeholder="Birthday ('day'/'month'/'year')"
                                required />
                        </div>

                        <div class="form-group cta-step">
                            <div class="cta prev">
                                <p class="cta-color">
                                    <span class="link_wrap">
                                        <a class="link_text" href="#" @click.prevent="prev()"><span
                                                class="arrow-prev"></span>Previous
                                        </a>
                                    </span>
                                </p>
                            </div>
                        </div>
                        <div class="register-btn">
                            <input type="submit" value="Créer mon compte" />
                        </div>
                    </form>
                </section>
            </transition>
        </section>
        <section class="congrats register" v-if="hasSeenCongrats">
            <div class="register-icon">
                <img class="register-icon-item" src="https://vuejs.org/images/logo.png" alt="vue logo">

            </div>
            <h2 class="congrats-title">Thank you !</h2>
            <p class="congrats-subtitle">
                You have successfully created your account and joined the<br />
                <strong>VueJS<br />multiple steps form</strong>
            </p>
        </section>
        <section class="pen-description">
            <p style="color: #fff;">by</p>
            <a href="https://www.behance.net/TMMcreation" target="_blank" class="pen-copyright">
                <img height="70" width="70" src="https://mir-s3-cdn-cf.behance.net/user/230/3855631.53be57a2b99fb.jpg"
                    alt="Digital strategist - Web designer">
            </a>
        </section>
    </div>
</template>